package controller

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/attapon-th/go-pkgs/zlog/log"
	"github.com/attapon-th/swagger-rapidoc/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var (
	minioClient   *minio.Client
	bucketSwagger = "swagger"
)

func EndpointSwagger(r fiber.Router) {
	ct, err := model.LoadCredential(viper.GetString("minioCredential"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration minio.")
	}
	// Initialize minio client object.
	minioClient, err = minio.New(ct.GetEndpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(ct.AccessKey, ct.SecretKey, ""),
		Secure: ct.UseSSL(),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect minio.")
	}

	r.Get("/", swaggerRapidoc)
	r.Get("/files/*", getFiles)
}

func getFiles(c *fiber.Ctx) error {
	filename := c.Params("*")
	if filename == "" {
		return fiber.ErrNotFound
	}
	log.Debug().Str("bucker", bucketSwagger).Str("filename", filename).Send()
	f, err := minioClient.GetObject(context.Background(), bucketSwagger, filename, minio.GetObjectOptions{})
	if err != nil {
		return fiber.ErrNotFound
	}
	info, err := f.Stat()
	if err != nil {
		log.Error().Err(err).Interface("info", info).Send()
		return err
	}
	b, _ := io.ReadAll(f)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.Send(b)
}

func swaggerRapidoc(c *fiber.Ctx) error {
	name := c.Query("name", "-")
	if name == "-" {
		return c.Status(404).SendString("Not found.")
	}

	q := c.Context().QueryArgs()
	q.Del("name")

	// theme		- light | dark
	// font-size	- default | large | largest
	// render-style	- read | view | focused
	rapidocAttributes := "theme|font-size|render-style"
	pathprefix := viper.GetString("app.prefix")
	attrs := fmt.Sprintf(` spec-url="%s/files/%s" `, pathprefix, name)
	for _, k := range strings.Split(rapidocAttributes, "|") {
		if v := c.Query(k, ""); v != "" {
			attrs += fmt.Sprintf(` %s="%s" `, k, v)
		}
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	title := strings.Title(strings.TrimSuffix(filepath.Base(name), filepath.Ext(name)))
	h := templateRapidoc(title, attrs)
	return c.SendString(h)
}

func templateRapidoc(s ...any) string {
	return fmt.Sprintf(`
<!doctype html> <!-- Important: must specify -->
<html>
	<head>
	<title>%s</title>
	<meta charset="utf-8"> <!-- Important: rapi-doc uses utf8 characters -->
	<script type="module" src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"></script>
	</head>
	<body>
	<rapi-doc %s"> </rapi-doc>
	</body>
</html>
`, s...)
}
