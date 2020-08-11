//
//  BeagleConfig.swift
//  ${project_name}
//
//  Created by ${organization_name} on ${date}.
//  Copyright © ${year} ${organization_name}. All rights reserved.
//

import Foundation
import Beagle

class BeagleConfig {

    static let dependencies = BeagleDependencies()

    /// Config Beagle
    static func config() {
        dependencies.urlBuilder = UrlBuilder(baseUrl: URL(string: Constants.serverAdress))
        Beagle.dependencies = dependencies

        BeagleConfig.registerWidgets()
        BeagleConfig.registerCustomActions()
    }

    /// Register custom widgets
    static private func registerWidgets() {

    }

    /// Register custom actions
    static private func registerCustomActions() {

    }
}
