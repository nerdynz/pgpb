<p align="center">
    <a href="https://pocketbase.io" target="_blank" rel="noopener">
        <img src="https://i.imgur.com/5qimnm5.png" alt="PocketBase - open source backend in 1 file" />
    </a>
</p>

## Fork Explanation  
[Pocketbase Documentation](https://pocketbase.io/docs)  
Pocketbase is a great product and very efficient for small-mid projects. It has no any additional setup for the database and it is very easy to use on a single server. But in our use-case we really need to use postgres as a main database and operate it manually. Also we love what Pocketbase does in CRUD operation and RBAC implementations so we want to use it. For these reasons we decided to fork it and make it compatible with postgres. We are still working on it and we will update the documentation as soon as we finish the project.  

## Usage  
You can easily fork and setup the project.  

```bash
# clone and download libraries
git clone https://github.com/AlperRehaYAZGAN/postgresbase
cd postgresbase
go mod download

# setup postgres, minio(s3) and mailhog(testing-mail-operations)
docker-compose up -d

# run the project with postgres connection info
CGO_ENABLED=0 go build -tags pq -o server \
    github.com/pocketbase/pocketbase/examples/base \
    env \
    LOGS_DATABASE="postgresql://user:pass@localhost/logs?sslmode=disable" \
    DATABASE="postgresql://user:pass@localhost/postgres?sslmode=disable" \
    ./server serve  

```
