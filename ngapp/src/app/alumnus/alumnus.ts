import { Component, signal, computed } from '@angular/core';
import { NgbPaginationModule } from '@ng-bootstrap/ng-bootstrap';
import { Header } from '../header/header';
import { Footer } from '../footer/footer';

@Component({
  selector: 'app-alumnus',
  standalone: true,
  imports: [NgbPaginationModule, Header, Footer],
  templateUrl: './alumnus.html',
  styleUrl: './alumnus.css',
})
export class Alumnus {
  searchText = signal('');
  items = signal([
    { name: 'John Doe', email: 'john@example.com' },
    { name: 'Sarah Smith', email: 'sarah@example.com' },
    { name: 'Michael Brown', email: 'michael@example.com' },
    { name: 'Alice Green', email: 'alice@example.com' },
    { name: 'David Wilson', email: 'david@example.com' },
    { name: 'Emma Taylor', email: 'emma@example.com' },
    { name: 'Doseok Kim', email: 'doseok@gmail.com' }
  ]);


  totalItems = 20; // Total number of items in the collection
  currentPage = 1; // The current active page, starts at 1
  itemsPerPage = 5; // Number of items per page
  page = signal(1); // current page

  filteredUsers = computed(() => {
    const text = this.searchText().toLowerCase();
    return this.items().filter((u: { name: string; email: string }) =>
      u.name.toLowerCase().includes(text) ||
      u.email.toLowerCase().includes(text)
    );
  });

  clearSearch() {
    this.searchText.set('');
    this.currentPage = 1;
  }

  loadPage(page: number) {
    // Logic to fetch data for the new page
    console.log('Page changed to:', page);
    // This is where you would typically update the displayed data,
    // either client-side (slicing an array) or server-side (making an API call with the new page number)
  }
}

