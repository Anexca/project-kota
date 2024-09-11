package sqlqueries

var ASSESSMENT_STATUS_QUERY string = "CREATE TYPE status AS ENUM ('COMPLETED', 'REJECETD', 'PENDING');"

var PAYMENT_STATUS_QUERY string = `CREATE TYPE payment_status AS ENUM (
    'CREATED',
    'AUTHORIZED',
    'CAPTURED',
    'FAILED',
    'REFUNDED',
    'PARTIALLY_REFUNDED',
    'PENDING',
    'PROCESSING',
    'CANCELLED',
    'DISPUTED'
);`
