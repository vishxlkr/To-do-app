ALTER TABLE todos ADD COLUMN user_id UUID NOT NULL


ALTER TABEL todos ADD CONTRAINTS fk_todos_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE

-- on delete cascade -> if the user is deleted , all the todos associated with him will also be deleted