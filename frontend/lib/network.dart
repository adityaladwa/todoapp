import 'dart:convert';

import 'package:http/http.dart' as http;

class Network {
  Future<LoginResponse> login(String email, String password) async {
    final request = new LoginRequest(email: email, password: password);
    final response = await http.post(Uri.http('localhost:8080', 'api/v1/users/login'),
        body: request.toJson());
    if (response.statusCode == 200) {
      return LoginResponse.fromJson(jsonDecode(response.body));
    } else {
      throw Exception("Failed to login ${response.body}");
    }
  }
}

class LoginRequest {
  final String email;
  final String password;

  LoginRequest({this.email, this.password});

  Map<String, dynamic> toJson() => {'email': email, 'password': password};
}

class LoginResponse {
  final String token;

  LoginResponse({this.token});

  factory LoginResponse.fromJson(Map<String, dynamic> json) {
    return LoginResponse(token: json['token']);
  }
}
