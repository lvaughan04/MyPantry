//
//  ContentView.swift
//  MyPantry
//
//  Created by Luke Vaughan on 11/12/24.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        VStack {
            Image(systemName: "globe")
                .imageScale(.large)
                .foregroundStyle(.tint)
            Text("Welcome to MyPantry")
                .font(.largeTitle)
                .padding()
            Text("Your personal kitchen inventory tracker")
                .padding()
        }
        .padding()
    }
}

#Preview {
    ContentView()
}
