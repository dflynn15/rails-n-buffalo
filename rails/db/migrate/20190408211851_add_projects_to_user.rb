class AddProjectsToUser < ActiveRecord::Migration[5.2]
  def change
    add_column :users, :project, :uuid
  end
end
