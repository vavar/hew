from typing import Mapping

from sqlalchemy import Column, Integer
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class Rating(Base):
    """
    A single user's rating of a single meal.
    """
    __tablename__ = 'rating'

    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, index=True, nullable=False)
    menu_id = Column(Integer, index=True, nullable=False)
    restaurant_id = Column(Integer, index=True, nullable=False)
    portion_size = Column(Integer)
    healthiness = Column(Integer)
    sweetness = Column(Integer)
    spice_level = Column(Integer)
    rating = Column(Integer)

    def __str__(self):
        return f'Rating of meal {self.menu_id} by user {self.user_id}: {self.rating}'

    def serialize(self):
        return {
            'id': self.id,
            'user_id': self.user_id,
            'menu_id': self.menu_id,
            'restaurant_id': self.restaurant_id,
            'ratings': {
                'portion_size': self.portion_size,
                'healthiness': self.healthiness,
                'sweetness': self.sweetness,
                'spice_level': self.spice_level,
                'rating': self.rating,
            },
        }

    @classmethod
    def create(cls, user_id: int, menu_id: int, restaurant_id: int, **ratings: Mapping[str, int]):
        """
        Add a rating for the attributes of this meal
        :param user_id: Id of the submitting user
        :param menu_id: Id of the meal
        :param restaurant_id: Id of the restaurant
        :param ratings: Map of ratings. Possible keys are:
            portion_size
            healthiness
            sweetness
            spice_level
            rating
        All keys are optional, and any provided should be integer scores 1-10
        """
        if any(1 > int(score) > 10 for score in ratings.values()):
            raise ValueError('Scores should be integers from 1-10 inclusive.')

        ALLOWED_ATTRS = frozenset(['portion_size', 'healthiness', 'sweetness', 'spice_level', 'rating'])
        invalid_attrs = [attr for attr in ratings.keys() if attr not in ALLOWED_ATTRS]
        if invalid_attrs:
            raise ValueError('Invalid attributes submitted: {}'.format(invalid_attrs))

        return cls(user_id=user_id, menu_id=menu_id, restaurant_id=restaurant_id, **ratings)