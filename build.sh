cd templates && yarn build 
cp -r templates/build/* ../templates_build/
cd .. && go build main.go -o app