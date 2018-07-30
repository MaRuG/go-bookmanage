# Revelでの書籍管理アプリケーション

# go get
go get github.com/revel/revel

# 起動方法
`revel run`
http://localhost:9000

# 機能
* SQLite3のDBを使用
    * なければ自動生成
* Google Books APIで書籍名・ISBNでの検索可能・登録可能
* 機能一覧
    * 書籍追加
    * 書籍検索
    * 登録書籍検索
