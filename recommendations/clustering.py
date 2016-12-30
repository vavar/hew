"""
This script partitions the responders to the lunch questionnaire into clusters with similar preferences
"""
from collections import defaultdict

import pandas as pd
from sklearn.cluster import KMeans

dataset = pd.read_csv('q_data.csv')
X = dataset.iloc[:, 2:].values
names = dataset.iloc[:, 1]

wcss = []
for i in range(1, len(X)):
    kmeans = KMeans(n_clusters=i, init='k-means++')
    kmeans.fit(X)
    wcss.append(kmeans.inertia_)
print(wcss)  # Pause here to select the number of clusters via "the elbow method".

# Fitting K-Means to the dataset using n_clusters determined above.
kmeans = KMeans(n_clusters=4, init='k-means++')
y_kmeans = kmeans.fit_predict(X)

groups = defaultdict(list)
for i, group in enumerate(y_kmeans):
    groups[group].append(names[i])

print(groups)
