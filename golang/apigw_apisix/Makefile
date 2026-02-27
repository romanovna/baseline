ADMIN_API = http://localhost:9180/apisix/admin
ADMIN_KEY = edd1c9f034335f136f87ad84b625c8f1
GATEWAY = http://localhost:9080

.PHONY: up down logs restart setup test lint

up:
	docker compose up --build -d

down:
	docker compose down -v

logs:
	docker compose logs -f

restart: down up

setup:
	@echo "=== Creating consumer: site ==="
	@curl -s -X PUT $(ADMIN_API)/consumers/site \
		-H "X-API-KEY: $(ADMIN_KEY)" \
		-H "Content-Type: application/json" \
		-d '{"username":"site","plugins":{"key-auth":{"key":"key_site"}}}' | jq .
	@echo ""
	@echo "=== Creating consumer: oms ==="
	@curl -s -X PUT $(ADMIN_API)/consumers/oms \
		-H "X-API-KEY: $(ADMIN_KEY)" \
		-H "Content-Type: application/json" \
		-d '{"username":"oms","plugins":{"key-auth":{"key":"key_oms"}}}' | jq .
	@echo ""
	@echo "=== Creating consumer: mobile ==="
	@curl -s -X PUT $(ADMIN_API)/consumers/mobile \
		-H "X-API-KEY: $(ADMIN_KEY)" \
		-H "Content-Type: application/json" \
		-d '{"username":"mobile","plugins":{"key-auth":{"key":"key_mobile"}}}' | jq .
	@echo ""
	@echo "=== Creating upstream: aggregator ==="
	@curl -s -X PUT $(ADMIN_API)/upstreams/1 \
		-H "X-API-KEY: $(ADMIN_KEY)" \
		-H "Content-Type: application/json" \
		-d '{"name":"aggregator","type":"roundrobin","nodes":{"aggregator:8081":1},"retries":2,"timeout":{"connect":5,"send":5,"read":5},"checks":{"active":{"type":"http","http_path":"/get-price?product_id=1&client_id=client_001","healthy":{"interval":5,"successes":2},"unhealthy":{"interval":5,"http_failures":3}}}}' | jq .
	@echo ""
	@echo "=== Creating route: /get-price ==="
	@curl -s -X PUT $(ADMIN_API)/routes/test \
		-H "X-API-KEY: $(ADMIN_KEY)" \
		-H "Content-Type: application/json" \
		-d '{"name":"get-price","uri":"/get-price","upstream_id":"1","plugins":{"key-auth":{}}}' | jq .
	@echo ""
	@echo "=== Setup complete ==="

test:
	@echo "=== Test 1: No auth → 401 ==="
	@curl -s -w "\nHTTP Status: %{http_code}\n" "$(GATEWAY)/get-price?product_id=1&client_id=client_001"
	@echo ""
	@echo "=== Test 2: site → client_001 (default), product 1 → promo 79.99 ==="
	@curl -s -H "apikey: key_site" "$(GATEWAY)/get-price?product_id=1&client_id=client_001" | jq .
	@echo ""
	@echo "=== Test 3: oms → client_002 (vip), product 1 → promo 69.99 ==="
	@curl -s -H "apikey: key_oms" "$(GATEWAY)/get-price?product_id=1&client_id=client_002" | jq .
	@echo ""
	@echo "=== Test 4: mobile → client_003 (wholesale), product 2 → regular 180.00 ==="
	@curl -s -H "apikey: key_mobile" "$(GATEWAY)/get-price?product_id=2&client_id=client_003" | jq .
	@echo ""
	@echo "=== Test 5: site → client_002 (vip), product 5 → promo 149.99 ==="
	@curl -s -H "apikey: key_site" "$(GATEWAY)/get-price?product_id=5&client_id=client_002" | jq .
	@echo ""
	@echo "=== Test 6: oms → client_001 (default), product 5 → regular 199.99 ==="
	@curl -s -H "apikey: key_oms" "$(GATEWAY)/get-price?product_id=5&client_id=client_001" | jq .
	@echo ""
	@echo "=== Test 7: Unknown product → 404 ==="
	@curl -s -H "apikey: key_site" "$(GATEWAY)/get-price?product_id=99&client_id=client_001" | jq .
	@echo ""
	@echo "=== All tests complete ==="

lint:
	@echo "No linter configured"
