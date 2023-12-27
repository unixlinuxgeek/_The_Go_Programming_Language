### 8.4.1 Небуферезированные каналы.

Операция отправления в небуферезированный канал блокирует go-подпрограмму до тех пор,
пока другая go-подпрограмма не выполнит соответствующее получение из того же канала,
после чего значение становится переданным, и обе go-  подпрограммы продолжаются.

И наоборот если первой сделанв попытка, выполнить операцию получения, принимающая go-подпрограмма
блокируется до тех пор, пока другая go-подпрограмма не выполнит отправление значения в тот же канал.

Связи по небуферизированному каналу приводит к синхронизации операции отправления и получения.
По этой причине небуферизированные каналы иногда называют синхронными.