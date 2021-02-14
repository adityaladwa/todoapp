import 'package:flutter/material.dart';

class LoginPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: <Widget>[
          Padding(
            padding: const EdgeInsets.all(20.0),
            child: TextFormField(
              cursorColor: Theme.of(context).cursorColor,
              decoration: InputDecoration(
                  labelText: "Email", border: OutlineInputBorder()),
            ),
          ),
          Padding(
            padding: EdgeInsets.all(20.0),
            child: TextFormField(
              cursorColor: Theme.of(context).cursorColor,
              decoration: InputDecoration(
                  labelText: "Password", border: OutlineInputBorder()),
            ),
          ),
          Padding(
            padding: EdgeInsets.all(20.0),
            child: ElevatedButton(
              onPressed: () {},
              child: Text("Login"),
            ),
          ),
        ],
      ),
    );
  }
}
