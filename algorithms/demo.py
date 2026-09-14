import numpy as np
import pandas as pd
from pathlib import Path
import sys

# Add root directory to path to import the algorithm
sys.path.append(str(Path(__file__).resolve().parents[1]))
from algorithms.Logistic_Regression.logistic_regression import LogisticRegressionScratch

# Load sample data
data_path = Path(__file__).parent / "sample_data.csv"
df = pd.read_csv(data_path)

X = df[['feature1', 'feature2']].values
y = df['target'].values

# Train model
model = LogisticRegressionScratch(learning_rate=0.1, num_iterations=1500)
model.fit(X, y)

# Make predictions
predictions = model.predict(X)
print("Predictions:", predictions)
print("Actual targets:", y)