class CreateFlags < ActiveRecord::Migration[5.2]
  def change
    create_table :flags, id: :uuid do |t|
      t.string :name
      t.boolean :enabled
    end
  end
end
