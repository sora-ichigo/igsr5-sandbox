# frozen_string_literal: true

require "bundler/inline"

gemfile(true) do
  source "https://rubygems.org"

  git_source(:github) { |repo| "https://github.com/#{repo}.git" }

  # Activate the gem you are reporting the issue against.
  gem "activerecord", "~> 7.0.0"
  gem "sqlite3"
end

require "active_record"
require "minitest/autorun"
require "logger"

# This connection will do for database-independent bug reports.
ActiveRecord::Base.establish_connection(adapter: "sqlite3", database: ":memory:")
ActiveRecord::Base.logger = Logger.new(STDOUT)

ActiveRecord::Schema.define do
  create_table :posts, force: true do |t|
    t.string :column_1
    t.string :column_2
  end
end

class Post < ActiveRecord::Base
  attr_accessor :column_1_changed, :column_2_changed
end

class BugTest < Minitest::Test
  def test_has_previously_changes
    post = Post.create!

    post.update(column_1: "first changed", column_2: "second changed")
    column_1_changed = post.column_1_previously_changed?
    column_2_changed = post.column_2_previously_changed?
    assert_equal(
      {column_1_changed: true, column_2_changed: true},
      {column_1_changed: column_1_changed, column_2_changed: column_2_changed}
    )
  end

  def test_has_only_last_previously_change
    post = Post.create!

    post.update(column_1: "first changed")
    post.update(column_2: "first changed")
    column_1_changed = post.column_1_previously_changed?
    column_2_changed = post.column_2_previously_changed?
    assert_equal(
      {column_1_changed: false, column_2_changed: true},
      {column_1_changed: column_1_changed, column_2_changed: column_2_changed}
    )
  end

  def test_has_only_last_saved_previously_change_in_transaction
    post = Post.create!

    ActiveRecord::Base.transaction do
      post.update(column_1: "first changed")
      post.update(column_2: "first changed")
    end

    column_1_changed = post.column_1_previously_changed?
    column_2_changed = post.column_2_previously_changed?
    assert_equal(
      {column_1_changed: false, column_2_changed: true},
      {column_1_changed: column_1_changed, column_2_changed: column_2_changed}
    )
  end
end
