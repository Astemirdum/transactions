
CREATE FUNCTION tv1.check_overdraft() RETURNS trigger AS $negative_balance$
BEGIN
    IF NEW.cash < 0 THEN
        RAISE EXCEPTION '% cannot have a negative balance', NEW.id;
    END IF;
    RETURN NULL;
END;
$negative_balance$ LANGUAGE plpgsql;

CREATE TRIGGER negative_balance
    AFTER UPDATE ON tv1.balances FOR EACH ROW
EXECUTE FUNCTION tv1.check_overdraft();


