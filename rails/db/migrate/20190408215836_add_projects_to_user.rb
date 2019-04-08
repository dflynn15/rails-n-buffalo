class AddProjectsToUser < ActiveRecord::Migration[5.2]
  def change
    add_column :users, :projects, :uuid
  end
end
