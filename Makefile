server:
	air

stylesmin:
	npx nodemon -e css -i pkg/themes/pigo.min.css -i pkg/themes/pigo.css \
		-x 'lightningcss --minify --bundle --targets ">=0.25%" pkg/themes/_input.css -o pkg/themes/pigo.min.css'

styles:
	npx nodemon -e css -i pkg/themes/pigo.min.css -i pkg/themes/pigo.css \
		-x 'lightningcss --bundle --targets ">=0.25%" pkg/themes/_input.css -o pkg/themes/pigo.css'

css:
	make -j2 stylesmin styles

deps:
	npm i && go install github.com/air-verse/air@latest && go mod tidy

dev:
	make -j3 stylesmin styles server
