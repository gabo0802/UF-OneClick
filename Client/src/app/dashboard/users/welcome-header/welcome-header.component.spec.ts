import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WelcomeHeaderComponent } from './welcome-header.component';

describe('WelcomeHeaderComponent', () => {
  let component: WelcomeHeaderComponent;
  let fixture: ComponentFixture<WelcomeHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        WelcomeHeaderComponent
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WelcomeHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('username should be initially empty', () => {
    expect(component.username).toEqual('');
  });

  it('currentDate should be today', () => {

    let todayDate: Date = new Date();

    expect(component.currentDate).toEqual(todayDate);
  });
});
