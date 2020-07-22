#!/usr/bin/python3
import json
from ml_template.notebook_template import Estimator, render_template


def Run(datapoints_amount, category, labeled_data):

    if datapoints_amount == "<50":
        print("You should get more data!")
        return

    if category == "false":
        algorithm = "regression"
    elif labeled_data == "true":
        algorithm = "classification"
    else:
        algorithm = "clustering"

    estimator = {
        "classification": {
            "<10k": Estimator.KNeighborsClassifier.value,
            "<100k": Estimator.RandomForestClassifier.value,
            ">100k": Estimator.SGDClassifier.value
        },
        "clustering": {
            "<10k": Estimator.MiniBatchKMeans.value,
            "<100k": Estimator.KMeans.value,
            ">100k": Estimator.KMeans.value
        },
        "regression": {
            "<10k": Estimator.SVR.value,
            "<100k": Estimator.SVR.value,
            ">100k": Estimator.SGDRegressor.value
        },
    }[algorithm][datapoints_amount]

    print(f"You should go for the {estimator} algorithm")
    print("Generating jupyter notebook...")
    filename = f"{estimator}.ipynb"

    with open(filename, 'w') as f:
        file_contents = render_template(estimator)
        f.write(json.dumps(file_contents))

    print(f"Done! please check {filename} and happy machine learning =]")
