templ:
	templ generate

tailwind:
	npx tailwindcss -i web/src/index.css -o web/public/styles.css --watch

.PHONY: templ tailwind