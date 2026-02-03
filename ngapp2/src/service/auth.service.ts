import { Injectable, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private apiUrl = 'https://goapp/api/login';

  user = signal<any>(null);

  constructor(private http: HttpClient) {}

  login(credentials: { email: string; password: string }) {
    return this.http.post(this.apiUrl, credentials).pipe(
      tap((res) => this.user.set(res))
    );
  }

  logout() {
    this.user.set(null);
  }
}