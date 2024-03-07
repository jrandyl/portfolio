templ:
	templ generate

tailwind:
	tailwindcss -i web/src/index.css -o web/public/styles.css --minify --watch

.PHONY: templ tailwind