import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UsernameFieldComponent } from './username-field.component';

describe('UsernameFieldComponent', () => {
  let component: UsernameFieldComponent;
  let fixture: ComponentFixture<UsernameFieldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UsernameFieldComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UsernameFieldComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
