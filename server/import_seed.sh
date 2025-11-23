#!/usr/bin/env sh

set -e

# Load environment variables from .env
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# Check required vars
if [ -z "$MONGO_URI" ] || [ -z "$MONGO_DB" ]; then
  echo "MONGO_URI or MONGO_DB not set in .env"
  exit 1
fi

shopt -s nullglob
JSON_FILES=(seed_data/*.json)

if [ ${#JSON_FILES[@]} -eq 0 ]; then
  echo "No JSON files found in seed_data/"
  exit 1
fi

echo "Starting import into database"

for FILE in "${JSON_FILES[@]}"; do
  COLLECTION=$(basename "$FILE" .json)

  echo "📥 Importing $FILE → collection: $COLLECTION"

  mongoimport \
    --uri "$MONGO_URI" \
    --db "$MONGO_DB" \
    --collection "$COLLECTION" \
    --jsonArray \
    --drop \
    --file "$FILE"
done

echo "✅ Import completed successfully."
