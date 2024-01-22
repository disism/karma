.PHONY: migrations

# authx atlas https://atlasgo.io/getting-started#installation
migrations:
	atlas migrate diff migration_name \
      --dir "file://ent/migrate/migrations" \
      --to "ent://ent/schema" \
      --dev-url "sqlite://file?mode=memory&_fk=1"
