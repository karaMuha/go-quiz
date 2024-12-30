CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS questions(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title TEXT NOT NULL,
  created_at timestamp(0) with time zone NOT NULL,
  updated_at timestamp(0) with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS question_options(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  question_id UUID NOT NULL,
  title TEXT NOT NULL,
  is_correct BOOLEAN NOT NULL,
  created_at timestamp(0) with time zone NOT NULL,
  updated_at timestamp(0) with time zone NOT NULL,

  FOREIGN KEY (question_id) REFERENCES questions (id) ON DELETE CASCADE,
  UNIQUE(question_id, title)
);