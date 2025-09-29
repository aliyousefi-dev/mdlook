# Quick Start Guide

1.  **Initialize Your Repository**
    Select your target folder and run:
    ```bash
    ovacli init .
    ```

---

2.  **Check Repository Status**
    To verify your setup, run:
    ```bash
    ovacli status
    ```

---

3.  **Index Your Videos**
    Scan and index all available videos in the selected folder:
    ```bash
    ovacli index
    ```
    This command will search for videos and add them to your database.

---

4.  **Optional: Optimize Videos (Cooking)**
    For improved performance and user experience, you can optimize your videos:
    ```bash
    ovacli cook
    ```
    > The cooking process generates story thumbnails, making them visible in the video player for easier navigation.

---

5.  **Start the Development Server**
    Launch the server with:
    ```bash
    ovacli serve
    ```