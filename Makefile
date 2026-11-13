.PHONY: dev-api dev-web dev

dev-api:
	cd api && air

dev-web:
	cd web && bun run dev

dev:
	make -j 2 dev-api dev-web
