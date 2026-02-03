import { Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AlumnusComponent } from './alumnus/alumnus.component';
import { GroupComponent } from './group/group.component';
import { ToolComponent } from './tool/tool.component';

export const routes: Routes = [
    { path: 'login', loadComponent: () => LoginComponent },
    { path: 'dashboard', component: DashboardComponent },
    { path: 'alumnus', component: AlumnusComponent },
    { path: 'group', component: GroupComponent },
    { path: 'tool', component: ToolComponent },
    { path: '', redirectTo: 'login', pathMatch: 'full' }
];
