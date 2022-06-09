
# validation-based protocol

In cases where a majority of transactions are read-only transactions, the `rate of conflicts among transactions` may be `low`. 

Thus, many of these transactions, if executed without the supervision of a concurrency-control scheme, would nevertheless leave the system in a consistent state.

Jadi, banyak dari transaksi ini, jika dijalankan tanpa pengawasan skema kontrol konkurensi, akan tetap meninggalkan sistem dalam keadaan konsisten.
WOW wkwk, pada skema optimis, kita yakin bahwa transaksi tanpa di kontrol pun akan tetap konsisten


This validation scheme is called the `optimistic concurrency-control` scheme since transactions execute optimistically, assuming they will be able to finish execution and validate at the end. 

In contrast, locking and timestamp ordering are pessimistic in that they force a wait or a rollback whenever a conflict is detected, even though there is a chance that the schedule may be conflict serializable.

# apa saja fasenya?
fase pada validation


