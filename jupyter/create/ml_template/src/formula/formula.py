#!/usr/bin/python3

import json
from formula.notebook_template import Estimator, render_template


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
            ">100k": Estimator.SGDClassifier.value,
        },
        "clustering": {
            "<10k": Estimator.MiniBatchKMeans.value,
            "<100k": Estimator.KMeans.value,
            ">100k": Estimator.KMeans.value,
        },
        "regression": {
            "<10k": Estimator.SVR.value,
            "<100k": Estimator.SVR.value,
            ">100k": Estimator.SGDRegressor.value,
        },
    }[algorithm][datapoints_amount]

    print("You should go for the {} algorithm".format(estimator))
    print("Generating jupyter notebook...")
    filename = "{}.ipynb".format(estimator)

    with open(filename, "w") as f:
        file_contents = render_template(estimator)
        f.write(json.dumps(file_contents))

    print(
        "Done! please check {} and happy machine \
    learning =]".format(
            filename
        )
    )
