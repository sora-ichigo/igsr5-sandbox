# frozen_string_literal: true

require "bundler/inline"

gemfile(true) do
  source "https://rubygems.org"

  git_source(:github) { |repo| "https://github.com/#{repo}.git" }

  # rails v7.0.6 を使いたい
  gem "rails", '7.0.6'
  gem "sqlite3"

  gem "paranoia", "2.6.2"
end

require "active_record"
require "minitest/autorun"
require "logger"

# This connection will do for database-independent bug reports.
ActiveRecord::Base.establish_connection(adapter: "sqlite3", database: ":memory:")
ActiveRecord::Base.logger = Logger.new(STDOUT)

ActiveRecord::Schema.define do
  create_table :posts, force: true do |t|
    t.string :title
  end

  create_table :comments, force: true do |t|
    t.string :title
    t.datetime :deleted_at
  end
end

class Post < ActiveRecord::Base
  after_destroy do
    self.title = "foo"
  end
end

class AfterDestroyTest < Minitest::Test
  def test_after_destroy
    post = Post.create!(title: "bar")

    assert_equal 1, Post.count
    assert_equal post.title, "bar"

    assert_raises(FrozenError) do
      post.destroy!
    end
  end
end

class Post < ActiveRecord::Base
  after_destroy_commit do
    self.title = "foo"
  end
end

class AfterDestroyTest < Minitest::Test
  def test_after_destroy
    post = Post.create!(title: "bar")

    assert_equal 1, Post.count
    assert_equal post.title, "bar"

    assert_raises(FrozenError) do
      post.destroy!
    end
  end
end

class Comment < ActiveRecord::Base
  after_destroy_commit do
    self.title = "foo"
  end

  acts_as_paranoid
end

class AfterDestroyWithParanoiaTest < Minitest::Test
  def test_after_destroy
    comment = Comment.create!(title: "bar")

    assert_equal 1, Comment.count
    assert_equal comment.title, "bar"

    # Failure:
    # AfterDestroyWithParanoiaTest#test_after_destroy [./activerecord_after_destroy_freeze.rb:89]:
    # FrozenError expected but nothing was raised.
    assert_raises(FrozenError) do # 通らない
      comment.destroy!
    end
  end
end