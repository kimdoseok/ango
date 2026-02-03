import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Alumnus } from './alumnus';

describe('Alumnus', () => {
  let component: Alumnus;
  let fixture: ComponentFixture<Alumnus>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Alumnus]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Alumnus);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
