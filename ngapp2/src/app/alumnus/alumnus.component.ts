import { Component, signal, computed } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NgxPaginationModule } from 'ngx-pagination';
import { HeaderComponent } from '../header/header.component';
import { FooterComponent } from '../footer/footer.component';

@Component({
  selector: 'app-alumnus',
  standalone: true,
  imports: [HeaderComponent, FooterComponent, CommonModule, FormsModule,NgxPaginationModule],
  templateUrl: './alumnus.component.html',
  styleUrl: './alumnus.component.scss'
})
export class AlumnusComponent {
  searchText = signal('');
  page = signal(1); // current page

  users = signal([
    { name: 'John Doe', email: 'john@example.com' },
    { name: 'Sarah Smith', email: 'sarah@example.com' },
    { name: 'Michael Brown', email: 'michael@example.com' },
    { name: 'Alice Green', email: 'alice@example.com' },
    { name: 'David Wilson', email: 'david@example.com' },
    { name: 'Emma Taylor', email: 'emma@example.com' },
    { name: 'Doseok Kim', email: 'doseok@gmail.com' }
  ]);

  filteredUsers = computed(() => {
    const text = this.searchText().toLowerCase();
    return this.users().filter(u =>
      u.name.toLowerCase().includes(text) ||
      u.email.toLowerCase().includes(text)
    );
  });

  clearSearch() {
    this.searchText.set('');
    this.page.set(1);
  }
}
/*
           
           [(ngModel)]="searchText()"
           (ngModelChange)="searchText.set($event); page.set(1)"
*/