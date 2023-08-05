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
end

class Post < ActiveRecord::Base
  after_commit do
    return
  end
end

# NOTE: after_commit でエラーが発生してしまうが、commit 完了時点で destroy は完了しているので以下のテストは通る
#
# # Running:
#
# D, [2023-07-19T08:00:33.711695 #6790] DEBUG -- :   Post Count (0.0ms)  SELECT COUNT(*) FROM "posts"
# D, [2023-07-19T08:00:33.712648 #6790] DEBUG -- :   TRANSACTION (0.0ms)  begin transaction
# D, [2023-07-19T08:00:33.712715 #6790] DEBUG -- :   Post Create (0.0ms)  INSERT INTO "posts" ("title") VALUES (?)  [["title", "bar"]]
# D, [2023-07-19T08:00:33.712769 #6790] DEBUG -- :   TRANSACTION (0.0ms)  commit transaction
# #<LocalJumpError: unexpected return>
# D, [2023-07-19T08:00:33.712919 #6790] DEBUG -- :   Post Count (0.0ms)  SELECT COUNT(*) FROM "posts"
# .
#
# Finished in 0.002236s, 447.2272 runs/s, 894.4545 assertions/s.
# 1 runs, 2 assertions, 0 failures, 0 errors, 0 skips
class AfterDestroyTest < Minitest::Test
  def test_after_destroy
    assert_equal 0, Post.count

    post = Post.create!(title: "bar")

    assert_equal 1, Post.count
  end
end

