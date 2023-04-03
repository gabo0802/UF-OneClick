import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddInactiveSubscriptionComponent } from './add-inactive-subscription.component';

describe('AddInactiveSubscriptionComponent', () => {
  let component: AddInactiveSubscriptionComponent;
  let fixture: ComponentFixture<AddInactiveSubscriptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddInactiveSubscriptionComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddInactiveSubscriptionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
