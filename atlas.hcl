variable "db_url" {
    type = string
    default = getenv("DATABASE_URL")
}

env "dev" {
    src = "file://schema.sql"
    url = var.db_url

    // Define the URL of the Dev Database for this environment
    // See: https://atlasgo.io/concepts/dev-database
    dev = "docker://postgres/15/dev?search_path=public"

    migration {
        dir = "file://migrations"
    }
}