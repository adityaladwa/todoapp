import 'package:flutter/material.dart';
import 'package:frontend/network.dart';

class LoginPage extends StatelessWidget {
  final network = new Network();

  final emailController = TextEditingController();
  final passwordController = TextEditingController();

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
              controller: emailController,
              decoration: InputDecoration(
                  labelText: "Email", border: OutlineInputBorder()),
            ),
          ),
          Padding(
            padding: EdgeInsets.all(20.0),
            child: TextFormField(
              cursorColor: Theme.of(context).cursorColor,
              controller: passwordController,
              decoration: InputDecoration(
                  labelText: "Password", border: OutlineInputBorder()),
            ),
          ),
          Padding(
            padding: EdgeInsets.all(20.0),
            child: ElevatedButton(
              onPressed: () {
                network
                    .login(emailController.text, passwordController.text)
                    .catchError((e) {
                  debugPrint("Error loging user");
                }).then((res) {
                  debugPrint(res.token);
                });
              },
              child: Text("Login"),
            ),
          ),
        ],
      ),
    );
  }
}
