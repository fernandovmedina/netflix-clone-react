DELIMITER //
CREATE PROCEDURE InsertPaymentAndGiftCode(
  IN plan_name VARCHAR(255),
  IN payment_date DATETIME,
  IN account_email VARCHAR(255),
  IN method_name VARCHAR(255),
  IN gift_code_value VARCHAR(255)
)
BEGIN
  DECLARE last_payment_id INT;
  INSERT INTO PAYMENTS (AMOUNT, PAYMENT_DATE, ID_ACCOUNT, ID_PAYMENT_METHOD)
  VALUES (
    (SELECT PRICE FROM PLANS WHERE NAME = plan_name LIMIT 1),
    payment_date,
    (SELECT ID_ACCOUNT FROM ACCOUNTS WHERE EMAIL = account_email LIMIT 1),
    (SELECT ID_PAYMENT_METHOD FROM PAYMENT_METHODS WHERE METHOD_NAME = method_name LIMIT 1)
  );
  SET last_payment_id = LAST_INSERT_ID();
  INSERT INTO GIFT_CODE_PAYMENTS (ID_CODE, ID_PAYMENT)
  VALUES (
    (SELECT ID_CODE FROM GIFT_CODES WHERE VALUE = gift_code_value LIMIT 1),
    last_payment_id
  );
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE InsertPaymentAndOxxo(
  IN plan_name VARCHAR(255),
  IN payment_date DATETIME,
  IN account_email VARCHAR(255),
  IN method_name VARCHAR(255),
  IN phone_number VARCHAR(255),
  IN client_name VARCHAR(255),
  IN oxxo_code VARCHAR(255)
)
BEGIN
  DECLARE last_payment_id INT;
  DECLARE last_oxxo_code INT;
  INSERT INTO PAYMENTS(AMOUNT,PAYMENT_DATE,ID_ACCOUNT,ID_PAYMENT_METHOD)
  VALUES (
    (SELECT PRICE FROM PLANS WHERE NAME = plan_name LIMIT 1),
    payment_date,
    (SELECT ID_ACCOUNT FROM ACCOUNTS WHERE EMAIL = account_email LIMIT 1),
    (SELECT ID_PAYMENT_METHOD FROM PAYMENT_METHODS WHERE METHOD_NAME = method_name LIMIT 1)
  );
  SET last_payment_id=LAST_INSERT_ID();
  INSERT INTO OXXO_CODES(VALUE,ID_ACCOUNT)VALUES(
    oxxo_code,
    (SELECT ID_ACCOUNT FROM ACCOUNT WHERE EMAIL=account_email)
  );
  SET last_oxxo_code=LAST_INSERT_ID();
  INSERT INTO OXXO_PAYMENTS(PHONE_NUMBER,NAME,ID_OXXO_CODE,ID_PAYMENT)
  VALUES (
    phone_number,
    client_name,
    last_oxxo_code,
    last_payment_id
  );
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE InsertPaymentAndCard(
  IN plan_name VARCHAR(255),
  IN payment_date DATETIME,
  IN account_email VARCHAR(255),
  IN method_name VARCHAR(255),
  IN card_number VARCHAR(255),
  IN card_due VARCHAR(255),
  IN card_cvv VARCHAR(255),
  IN card_name VARCHAR(255)
)
BEGIN
  DECLARE last_payment_id INT;
  INSERT INTO PAYMENTS(AMOUNT,PAYMENT_DATE,ID_ACCOUNT,ID_PAYMENT_METHOD)
  VALUES (
    (SELECT PRICE FROM PLANS WHERE NAME = plan_name LIMIT 1),
    payment_date,
    (SELECT ID_ACCOUNT FROM ACCOUNTS WHERE EMAIL = account_email LIMIT 1),
    (SELECT ID_PAYMENT_METHOD FROM PAYMENT_METHODS WHERE METHOD_NAME = method_name LIMIT 1)
  );
  SET last_payment_id = LAST_INSERT_ID();
  INSERT INTO CARD_PAYMENTS(CARD_NUMBER,CARD_DUE_DATE,CARD_CVV,ID_PAYMENT,CARD_NAME)
  VALUES (
    card_number,
    card_due,
    card_cvv,
    last_payment_id,
    card_name
  );
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE InsertOxxoCode(
  IN p_value TEXT,
  IN p_source TEXT,
  IN p_email VARCHAR(255),
  IN p_counter INT
)
BEGIN
  DECLARE last_id INT;
  INSERT INTO OXXO_CODES(VALUE,SOURCE,EMAIL,COUNTER)VALUES(p_value,p_source,p_email,p_counter);
  SET last_id=LAST_INSERT_ID();
  SELECT VALUE, SOURCE FROM OXXO_CODES WHERE ID_OXXO_CODE=last_id LIMIT 1;
END //
DELIMITER ;
