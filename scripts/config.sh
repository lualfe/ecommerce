cd config
touch test.yaml
echo 'PG_CONNECT: "host=localhost port=9400 dbname=ecommercetest user=ecommercetest password=abc@123 application_name=ecommercetest sslmode=disable"' >> default.yaml
echo 'GOOGLE_API_KEY: "AIzaSyBD06-SxpU2-S_nY0YisjDQKcCiRRPlwHQ"' >> default.yaml
echo 'jwt_key: "abc@123"' >> default.yaml