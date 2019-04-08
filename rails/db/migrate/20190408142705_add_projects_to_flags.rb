class AddProjectsToFlags < ActiveRecord::Migration[5.2]
  def change
    add_column :flags, :project_id, :uuid
  end
end
