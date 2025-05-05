CREATE TABLE appointments(
    appointment_id BIGSERIAL PRIMARY KEY,
    patient_id     INT NOT NULL,
    staff_id       INT,
    status_        VARCHAR(20) NOT NULL,
    created_at     TIMESTAMP DEFAULT now(),
    assigned_by    INT,

    FOREIGN KEY (patient_id)  REFERENCES users(user_id),
    FOREIGN KEY (staff_id)    REFERENCES users(user_id),
    FOREIGN KEY (assigned_by) REFERENCES users(user_id)
);