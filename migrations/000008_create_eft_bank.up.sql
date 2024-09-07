BEGIN;

CREATE TABLE bank (
    transactionid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transactiondate TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    amount INT NOT NULL,
    balance INT NOT NULL,
    signings VARCHAR(50),
    transactiontype VARCHAR(50) NOT NULL
    );


COMMIT;