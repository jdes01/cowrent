DDD + Clean Code + Clean Architecture + Hexagonal Architecture Dockerized Go and Nextjs Monorepo.

Using localstack, postgres, grafana+loki+prometheus.

compile-contracts.sh allows me to compile protobufs contracts, placing them in {web/api}/contracts/{dtos/request/responses}, working as contract libraries.

Gin and gorm are also used in the back.

I made my own rust-lookalike result micro library. You can find it in /api/utils/result
I also made my own loggin system, super easy to use. You can find it in /api/utils/logger
