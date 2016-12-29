from typing import Mapping


class Rating:
    """
    A single user's rating of a single meal.
    """

    def __init__(self, user_id, menu_id, **kwargs):
        self.user_id = user_id
        self.menu_id = menu_id
        self.portion_size = kwargs.get('portion_size')
        self.healthiness = kwargs.get('healthiness')
        self.sweetness = kwargs.get('sweetness')
        self.spice_level = kwargs.get('spice_level')
        self.rating = kwargs.get('rating')


def rate_menu(user_id: int, menu_id: int, attrs: Mapping[str, int]):
    """
    Add a rating for the attributes of this meal
    :param user_id: Id of the submitting user
    :param menu_id: Id of the meal
    :param attrs: Map of attributes. Possible attributes are:
        portion_size
        healthiness
        sweetness
        spice_level
        rating
    All attributes are optional, and all should be integer scores 1-10
    """
    if any(1 < int(score) > 10 for score in attrs.values()):
        raise ValueError('Scores should be integers from 1-10 inclusive.')

    ALLOWED_ATTRS = frozenset(['portion_size', 'healthiness', 'sweetness', 'spice_level', 'rating'])
    invalid_attrs = [attr for attr in attrs.keys() if attr not in ALLOWED_ATTRS]
    if invalid_attrs:
        raise ValueError('Invalid attributes submitted: {}'.format(invalid_attrs))

    rating = Rating(user_id, menu_id, **attrs)
