# rappel conso:

Simple bot scraping datagouv's "rappel conso" API to dispatch alerts when new product are announced. 

# does:

1. fetch records
2. filter them to only get last ones not recorded
3. if no new records: exit and log
4. insert new records
5. send post: twitter & facebook


# next:

- Migration
- Github flow
- Count routes
- UI with elm
- alerts geoloc
- alerts by email
- alerts by phone
