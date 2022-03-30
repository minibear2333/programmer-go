pg = db.getSiblingDB('pg')

pg.createCollection("comments")
pg.createCollection("hard_status")
pg.createCollection("interviews")
pg.createCollection("interviews_tags")
pg.createCollection("message_center")
pg.createCollection("message_config")
pg.createCollection("user")