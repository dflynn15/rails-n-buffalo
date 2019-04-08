class CreateProjects < ActiveRecord::Migration[5.2]
  def change
    create_table :projects, id: :uuid do |t|
      t.string :name
      t.uuid :flags
    end
  end
end
