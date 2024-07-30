# test challenge with PoW

# 1. Run solution

```
docker-compose up -d
```

# 2. Select of PoW algorithm

For this solution variation of Hashcash algorithm was selected. Main factors for this decision:

- a lot of documentation in internet
- easy to implement
- big difference between client and server work with increasing difficult

# 3. Architect notes

- challenge-response protocol was split for two request. Keep alive tcp connect while client calculate hash can be problematic and potential risk for new DOS attack.
- for remembering what task already sended used Redis as ready to use, high performance solution.
- project split to two part: one for client, second for backend. Usually client and server developed by different teams and on different stack. So I split them to two undepemdent program.