class Flag < ApplicationRecord
  belongs_to :project

  validates_presence_of :enabled
  validates_presence_of :name
end
