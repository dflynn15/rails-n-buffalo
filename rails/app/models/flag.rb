class Flag < ApplicationRecord
  belongs_to :project

  validates_inclusion_of :enabled, :in => [true, false]
  validates_presence_of :name
end
