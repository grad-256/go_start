PostgreSQLのインストール

brew install postgresql

文字コードをUTF-8でデータベースの初期化

initdb /usr/local/var/postgres -E utf8
initdb /opt/homebrew/var/postgresql@14 -E utf8

Postgresのバージョンチェック

postgres --version

PostgreSQLサーバの起動

postgres -D /usr/local/var/postgres

psql -lでデータベース一覧を取得

psql -l

