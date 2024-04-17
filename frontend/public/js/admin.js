/*// JavaScript for admin pages

// Function to handle admin login form submission
function handleAdminLogin() {
    var adminId = document.getElementById('adminId').value;
    var password = document.getElementById('password').value;

    // Perform client-side validation if needed

    // Send login request to backend API
    fetch('/api/admin/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            admin_id: adminId,
            password: password
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to login');
        }
        return response.json();
    })
    .then(data => {
        // Redirect to admin home page or display success message
        console.log('Admin login successful:', data);
        window.location.href = '/admin/home.html';
    })
    .catch(error => {
        // Display error message to user
        console.error('Admin login failed:', error);
        alert('Admin login failed. Please try again.');
    });
}

// Function to handle navigation to admin home page
function navigateToAdminHome() {
    window.location.href = '/admin/home.html';
}

// Function to handle navigation to admin profile page
function navigateToAdminProfile() {
    window.location.href = '/admin/profile.html';
}

// Function to handle navigation to admin marks page
function navigateToAdminMarks() {
    window.location.href = '/admin/marks.html';
}

// Function to handle navigation to admin attendance page
function navigateToAdminAttendance() {
    window.location.href = '/admin/attendance.html';
}

// Function to handle adding marks
function addMarks() {
    // Implement functionality to add marks
    console.log('Functionality to add marks');
}

// Function to handle adding attendance
function addAttendance() {
    // Implement functionality to add attendance
    console.log('Functionality to add attendance');
}*/

// admin.js login

// Function to handle admin login form submission
function handleAdminLogin() {
    var adminId = document.getElementById('adminId').value;
    var password = document.getElementById('password').value;

    // Perform client-side validation if needed

    // Send login request to backend API
    fetch('/api/admin/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            admin_id: adminId,
            password: password
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to login');
        }
        return response.json();
    })
    .then(data => {
        // Redirect to admin home page or display success message
        console.log('Admin login successful:', data);
        window.location.href = '/admin/home.html';
    })
    .catch(error => {
        // Display error message to user
        console.error('Admin login failed:', error);
        alert('Admin login failed. Please try again.');
    });
}
