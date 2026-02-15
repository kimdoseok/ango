import { Routes } from '@angular/router';
import { Login } from './login/login';
import { Dashboard } from './dashboard/dashboard';
import { Alumnus } from './alumnus/alumnus';
import { Group } from './group/group';
import { Tool } from './tool/tool';

export const routes: Routes = [
    { path: 'login', component: Login },
    { path: 'dashboard', component: Dashboard },
    { path: 'alumnus', component: Alumnus },
    { path: 'group', component: Group },
    { path: 'tool', component: Tool },
    { path: '', redirectTo: 'login', pathMatch: 'full' }
];

