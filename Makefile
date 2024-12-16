server:
	air

css:
	npx nodemon -e css -i pkg/themes/output.css -x \
		'lightningcss --minify --bundle --targets ">=0.25%" pkg/themes/input.css -o pkg/themes/output.css'

deps:
	npm i && go install github.com/air-verse/air@latest && go mod tidy

dev:
	make -j2 css server
